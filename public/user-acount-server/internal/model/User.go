package model

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

type User struct {
	
	Id int64 			`json:"id"` //数据库自增id
	Name string 		`json:"name"` //用户昵称
	UserRealName string `json:"user_real_name"`//用户真实姓名
	IdNo string 		`json:"id_no"` //用户唯一id
    Sex string			`json:"sex"` //用户性别
    Mobile string 		`json:"mobile"` //手机号
    Address string 		`json:"address"` //地址
	CreateAt string 	`json:"create_at"`//创建时间
	CreateIp string 	`json:"create_ip"`//用户创建的ip
	CreatBy string 		`json:"creat_by"`//创建人 sys

}


/**
用户curd详细信息
 */
type UserDetail struct {

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







 
 
 

