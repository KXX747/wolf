package model


const (
	CRAD_ONE = "crad_id_frist_img"
	CRAD_TWO = "crad_id_secode_img"


)

/**
用户curd用户信息
 */
/**
CREATE TABLE `union_user` (
  `id` bigint(20) NOT NULL AUTO_INCREMENT,
  `name` varchar(50) NOT NULL,
  `id_no` varchar(18) NOT NULL,
  `sex`  varchar(2) NOT NULL,
  `mobile` varchar(11) NOT NULL,
  `address` varchar(150) NOT NULL,
  `create_time` datetime DEFAULT NULL,
  `create_by` varchar(50) DEFAULT NULL,
  `create_by` varchar(50) DEFAULT NULL,
  PRIMARY KEY (`id`),
  KEY `index_unionuser_name_idno_union_no` (`name`,`id_no`,`mobile`)
) ENGINE=InnoDB AUTO_INCREMENT=100778 DEFAULT CHARSET=utf8;
 */

 //用户信息
type User struct {
	Id int64 			`json:"id"` //数据库自增id
	Name string 		`json:"name"` //用户昵称
	IdNo string 		`json:"id_no"` //用户唯一id
    Mobile string 		`json:"mobile"` //手机号
    Address string 		`json:"address"` //地址
	CreateAt string 	`json:"create_at"`//创建时间
	CreateIp string 	`json:"create_ip"`//用户创建的ip
	CreatBy string 		`json:"creat_by"`//创建人 sys
}


/**
用户curd详细信息
 */
type UserCommon struct {
	Id int64 				`json:"id"` //数据库自增id
	IdNo string 			`json:"id_no"`
	UserRealName string 	`json:"user_real_name"`//用户真实姓名
	CradId string 			`json:"crad_id"`//身份证
	CradIdFristImg string 	`json:"crad_id_frist_img"`//身份证正面
	CradIdSecodeImg string 	`json:"crad_id_secode_img"`//身份证正面
	Age string 				`json:"age"`//年龄
	Sex string				`json:"sex"` //用户性别
	Mobile string 			`json:"mobile"` //手机号
	Address string 			`json:"address"` //地址

}

//用户登录返回用户信息+token
type LoginResponse struct {
	Name string 		`json:"name"` //用户昵称
	IdNo string 		`json:"id_no"` //用户唯一id
	Mobile string 		`json:"mobile"` //手机号
	Address string 		`json:"address"` //地址
	CreateAt string 	`json:"create_at"`//创建时间
	CreateIp string 	`json:"create_ip"`//用户创建的ip
	CreatBy string 		`json:"creat_by"`//创建人 sys
	Token string 		`json:"token"`
}


//add upload  user
type ParamUpload struct {
	IdNo string 			`form:"id_no"  validate:"gt=0,required"`
	UserRealName string 	`form:"user_real_name"  validate:"gt=0,required"`//用户真实姓名
	CradId string 			`form:"card_id"  validate:"gt=0,required"`//身份证
	//CradIdFristImg string   `form:"crad_id_frist_img" validate:"required"`
	//CradIdSecodeImg string  `form:"crad_id_secode_img" validate:"required"`
	Age string 				`form:"age" validate:"gt=0,required"`
	Sex string 				`form:"sex" validate:"gt=0,required"`
}

//登录用户参数
type LoginInSystem struct {
	Name string 		`form:"name"  validate:"gt=0,required"` //用户昵称
	Mobile string 		`form:"mobile"  validate:"gt=0,required"` //手机号
}

//推出用户登录参数
type LoginOutSystem struct {
	Token string 		`form:"token"  validate:"gt=0,required"`
	Name string 		`form:"name"  validate:"gt=0,required"` //用户昵称
	Mobile string 		`form:"mobile"  validate:"gt=0,required"` //手机号
}





 
 
 

