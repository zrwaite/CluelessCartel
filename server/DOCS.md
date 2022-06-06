# Clueless Cartel API Docs

### **By default, all requests take username (and dev automatically) as parameters.**

### Default HTTP Method: POST; The function param should be included in any endpoint with :\<functionName>

---

| Endpoint       | Params                               | Response | Description                                                                         |
| -------------- | ------------------------------------ | -------- | ----------------------------------------------------------------------------------- |
| GET: /api/user |                                      | USER     | Gets all user data                                                                  |
| /api/user      | password: string                     | USER     | Creates a new user                                                                  |
| /api/base :new | location: [Location](#location-type) | USER     | If the user has enough cash, it will create a new base and return the updated user. |

---

### Location Type:

{"Nevada", "Florida", "Canadian Border", "New York"}

Operations:

Instant:

-   Buy Base
-   Buy chemicals
-   Buy hexagon

Timed/Instant Claim

-   Make drugs
-   Sell drugs
-   Buy structure for hexagon
-   Upgrade structure for hexagon

Probabilistic:

-   Rob drugs
-   Drive by
-   Blow up enemy structure
