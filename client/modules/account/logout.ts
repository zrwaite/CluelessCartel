const logout = () => {
	Cookies.delete("username")
	Cookies.delete("token")
	alert("Logged out")
}