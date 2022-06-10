const findUser = async ():Promise<boolean> => {
	const username = Cookies.get('username')
	if (!username) return false
	let apiRes = await httpReq("/api/user?username="+username)
	if (!apiRes) throw Error("Something went wrong")
	if (apiRes.Success) return true
	else return false
}

const forceMoveUser = async (from: 'login'|'game') => {
	const userFound = await findUser()
	if (userFound && from==='login') {
		forceNavigate('/game')
	} else if (!userFound && from==='game') {
		forceNavigate('/login')
	}
}