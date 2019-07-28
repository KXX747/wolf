package liveconsul


type K interface {
	//
	StoreKeyValue()(string);
	GetKeyValue()(value []byte,err error);

}


