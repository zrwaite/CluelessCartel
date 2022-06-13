const frontendUrl = DEV?'/pages':'/client/pages'
const forceNavigate = (to: string) => {
	window.location.href = frontendUrl + to
}
