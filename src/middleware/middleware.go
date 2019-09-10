package middleware

import (
	"encoding/json"
	"github.com/thedevsaddam/govalidator"
	"net/http"
)

func MyMiddleware(next http.Handler) http.Handler {
    return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
	rules := govalidator.MapData{
	    "id":			[]string{"required"},
	    "name":			[]string{"required", "between:1,255"},
	    "mimetype":			[]string{"required"},
	    "size":			[]string{"required"},
	    "file:file":		[]string{"required"},
	}

	messages := govalidator.MapData{
	    "name":			[]string{"required: Field \"Name\" is required", "between: Field \"Name\" length from 1 to 255"},
	}

	opts := govalidator.Options{
	    Request:		r,        // request object
	    Rules:		rules,    // rules map
	    Messages:		messages, // custom message map (Optional)
	    RequiredDefault:	true,     // all the field to be pass the rules
	}

	v := govalidator.New(opts)
	e := v.Validate()

	if len(e) > 0 {
	    err := map[string]interface{}{"error": e}
	    w.Header().Set("Content-type", "application/json")
	    json.NewEncoder(w).Encode(err)
	} else {
	    next.ServeHTTP(w, r)
	}
    })
}