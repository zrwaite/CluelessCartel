const findUser = async ():Promise<boolean> => {
	const username = Cookies.get('username')
	if (!username) return false
	let apiRes = await httpReq("/api/user?username="+username)
	if (!apiRes) throw Error("Something went wrong")
	if (apiRes.Success) return true
	else return false
}

const forceMoveUser = async () => {
	if (await findUser()) {
		forceNavigate('/game')
	} else {
		forceNavigate('/login')
	}
}