const frontendUrl = DEV?'/client/pages':''
const forceNavigate = (to: string) => {
	window.location.href = frontendUrl + to
}
