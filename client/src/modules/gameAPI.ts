interface APIResponse {
    Status: number,
    Success: boolean,
    Errors: string[]
    Response: any
}

export const gameAPI = async (url:string, method:string = "GET", params:any = {}):Promise<APIResponse | false> => {
    url = baseURL + "/api" + url;
    if (method !== "GET" && method !== "POST" && method !== "PUT" && method !== "DELETE") {
        console.error("invalid method");
        return false;
    }
    let headers:any = {}
    if (Cookies.get("token")) headers = {
        Authorization: "Bearer "+ Cookies.get("token")
    }
    try {
        let response;
        if (method === "GET") {
            response = await fetch(url, {
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
                body: JSON.stringify(params) // body data type must match "Content-Type" header
            });
        }
        const data = await response.json() as APIResponse;
        if (!response.ok) {
            return Promise.resolve(data);
        }
        return Promise.resolve(data);
    } catch (error) {
        console.error(error);
        return Promise.resolve(false);
        // return Promise.reject(false);
    }
}