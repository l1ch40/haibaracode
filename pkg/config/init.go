package config

func init() {
	Server = (&server{}).Load("conf/app.ini").Init()
	MySQL = (&mysql{}).Load("conf/db.ini").Init()
	Cookie = (&cookie{}).Load("conf/app.ini").Init()
}
