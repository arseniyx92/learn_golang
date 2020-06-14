## **Redirect**

`http.Redirect(w, req, "/", http.StatusSeeOther)`

_StatusMovedPermanently = **301**_ - save cash in your browser and always redirect to the same thing

_StatusSeeOther         = **303**_ - redirect to **GET**

_StatusTemporaryRedirect = **307**_ - stay on the **same** method, as it passed


#### 1 variant

`w.Header().Set("Location", "/")`

`w.WriteHeader(http.StatusSeeOther)`

#### 2 variant

`http.Redirect(w, req, "/", http.StatusSeeOther)`

`http.Redirect(w, req, "/", http.StatusTemporaryRedirect)`

`http.Redirect(w, req, "/", http.StatusMovedPermanently)`