const baseURL = DEV?"http://localhost:8004":"https://clueless-cartel.herokuapp.com";

interface APIResponse {
    Status: number,
    Success: boolean,
    Errors: string[]
    Response: any
}

const httpReq = async (url:string, method:string = "GET", params:any = {}):Promise<APIResponse | false> => {
    url = baseURL + url;
    if (method !== "GET" && method !== "POST" && method !== "PUT" && method !== "DELETE") {
        console.error("invalid method");
        return false;
    }
    let headers:any = {}
    if (Cookies.get("token")) headers = {
        // Authorization: "Bearer "+ Cookies.get("token")
    }
    try {
        let response;
        if (method === "GET") {
            response = await fetch(addQueryParam(url, 'dev', DEV?"true":"false"), {
                cache: 'no-cache',
                headers: headers
            });
        } else {
            response = await fetch(url, {
                method: method, // *GET, POST, PUT, DELETE, etc.
                cache: 'no-cache',
                // mode: 'no-cors',
                headers: {
                    ...headers,
                    'Content-Type': 'application/json',
                },
                body: JSON.stringify({...params, dev:DEV?"true":"false"}) // body data type must match "Content-Type" header
            });
        }
        const data = await response.json() as APIResponse;
        return Promise.resolve(data);
    } catch (error) {
        console.error(error);
        return Promise.reject(JSON.stringify(error));
    }
}