const tryLogin = async (usernameInputId: string, passwordInputId: string) => {
	const {username, password} = tryGetAccountInput(usernameInputId, passwordInputId)
	let apiRes = await httpReq("/api/signin", "POST", {
		username,
		password
	})
	if (!apiRes) throw Error("Something went wrong")
	if (apiRes.Success) {
		alert("Signed In")
		Cookies.set('username', username)
		Cookies.set('token', apiRes.Response.Token)
		forceNavigate('/game')
	} else {
		alert(JSON.stringify(apiRes.Errors))
	}
}

const trySignUp = async (usernameInputId: string, passwordInputId: string) => {
	const {username, password} = tryGetAccountInput(usernameInputId, passwordInputId)
	let apiRes = await httpReq("/api/user", "POST", {
		username,
		password
	})
	if (!apiRes) throw Error("Something went wrong")
	if (apiRes.Success) {
		alert("User created")
		Cookies.set('username', username)
		Cookies.set('token', apiRes.Response.Token)
		forceNavigate('/game')
	} else {
		alert(JSON.stringify(apiRes.Errors))
	}
}

const tryGetAccountInput = (usernameInputId: string, passwordInputId: string): {username: string, password: string} => {
	const usernameInput:HTMLInputElement|null = document.querySelector(`#${usernameInputId}`)
	const passwordInput:HTMLInputElement|null = document.querySelector(`#${passwordInputId}`)
	let [username, password] = ["",""];
	if (usernameInput) username = usernameInput.value;
	else throw Error("Failed to get #" + usernameInputId)
	if (passwordInput) password = passwordInput.value;
	else throw Error("Failed to get #" + passwordInputId)
	return {username, password}
}