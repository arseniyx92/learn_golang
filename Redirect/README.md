**Redirect**

`http.Redirect(w, req, "/", http.StatusSeeOther)`

_StatusMovedPermanently = 301_ - save cash in your browser and always redirect to the same thing

_StatusSeeOther         = 303_ - redirect to **GET**

_StatusTemporaryRedirect = 307_ - stay on the **same** method, as it passed