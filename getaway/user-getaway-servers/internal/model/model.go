package model

// Kratos hello kratos.
type Kratos struct {
	Hello string
}


//add user 
type ParamAddUser struct {

	Name string 	`form:"name"  validate:"gt=0,required"`
	Mobile string	`form:"mobile"  validate:"gt=0,required"`
}

//delete user
type ParamDeleteUser struct {

	IdNo string 	`form:"id_no"  validate:"gt=0,required"`
	Content string	`form:"content"  validate:"gt=0,required"`
}

//delete user
type ParamUpdateUser struct {

	IdNo string 	`form:"id_no"  validate:"gt=0,required"`
	Name string 	`form:"name"  validate:"gt=0,required"`
	Mobile string	`form:"mobile"  validate:"gt=0,required"`
	Address string  `form:"address"  validate:"gt=0,required"`
}
