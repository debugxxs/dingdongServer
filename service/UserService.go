package service

import (
	"dindongwork/common"
	"dindongwork/db"
	"dindongwork/models"
)

type UserService struct {
	db.UserDb
}


//CheckUserPass 验证用户名和密码
func (us UserService)CheckUserPass(userName,userPass string)bool{
	/*
	1.将密码解密
	2.传用户名给数据层查询
	3.返回查询的结果
	*/
	passWord := common.DecodePass(userPass)

	user := us.QueryUserPass(userName)
	DataPassWord := common.DecodePass(user.PassWord)
	if user.UserName == userName && DataPassWord == passWord {
		return true
	}
	return false
}

//CheckUserDataFunc 返回登陆结果
func (us UserService)CheckUserDataFunc(userName string)interface{}{
	/*
	1.根据用户名查询用户角色
	2.组织map返回
	*/
	userRole := us.QueryUserRole(userName)
	userData := map[string]interface{}{
		"userId":userRole.UserId,
		"userName":userRole.UserName,
		"userRole":userRole.RoleName,
		"userPhone":userRole.Phone,
		"userEmail":userRole.Email,
		"userAvatar":userRole.Avatar,
		"userPosition":userRole.Position,
	}
	return userData

}

//SaveUserToken 保存token到数据表
func (us UserService)SaveUserToken(userName ,token string){
	/*
	将用户token保存到数据库表格
	*/
	us.SaveToken(userName,token)
}

//CheckUser 检查用户密码
func (us UserService)CheckUser(userName ,token,userPass string) (string,bool){
	/*
	1.验证token
	2.修改用户名密码
	*/

	_,users:=us.QueryUserToken(userName,token)
	//fmt.Println(users)
	if users.Token == token{
		//用户验证安全，修改密码
		msg,result:=us.ModifyPassword(userName,userPass)
		if result !=0{
			return msg,true
		}else {
			return common.UpDataFail,true
		}
	}
	return common.ModifyPassErr,false
}

//GetListData 获取所有好友列表
func (us UserService)GetListData()(string,bool,interface{}){
	//1 获取组织架构下的组织列表
	//1.1 先根据组织用户来查询到相同的用户，并把他们放到同一个列表
	//2 根据组织id查询列表的值，并根据层级结构设置key
	//3 组织数据并返回
	//组织三层结构

	data := make([]interface{},0)
	msg,res,organizations :=us.GetAllOrganization()
	if res{
		for _,v := range organizations{
			//根据id去查询对应的用户列表
			_,_,userList :=us.QueryUserOrganization(v.OrganizationId)
			if v.ThreeLayer ==""{
				topData := make(map[string]interface{})
				twoData := make(map[string]interface{})
				userNames := make([]models.UserListData,0)
				if v.TwoLayer == ""{
					for _,lv := range userList{
						users := models.UserListData{UserId: lv.UserId,UserName: lv.UserName}
						userNames = append(userNames,users)
					}
					topData[v.TOPLayer] = userNames
					data = append(data,topData)
				}else {
					for _,lv := range userList{
						users := models.UserListData{UserId: lv.UserId,UserName: lv.UserName}
						userNames = append(userNames,users)
					}
					twoData[v.TwoLayer] = userNames
					topData[v.TOPLayer] = twoData
					data = append(data,topData)
				}
			}else {

				userNames := make([]models.UserListData,0)
				topData := make(map[string]interface{})
				twoData := make(map[string]interface{})
				threeData := make(map[string]interface{})
				for _,lv :=range userList{
					users := models.UserListData{UserId: lv.UserId,UserName: lv.UserName}
					userNames = append(userNames,users)
				}
				threeData[v.ThreeLayer] = userNames
				twoData[v.TwoLayer] = threeData
				topData[v.TOPLayer] = twoData

				data = append(data,topData)
			}

		}
		return msg,res,data
		//fmt.Println("输出结果：",data)
	}else {
		return msg,res,data
	}
}
