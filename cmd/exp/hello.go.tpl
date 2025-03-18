<h1>Hello {{.Name}}</h1>
<p>Bio {{.Bio}}</p>
{{if eq .Name "John Dough"}}
<p>This guy sucks</p>
{{else}}
<p>This guy is alright</p>
{{end}}
<h2>Friends of {{.Name}} </h2>
{{range .Friends}}
<p>{{ . }} </p>
{{end}}


<script>
    const user = {
        name: {{.Name}},
        bio: {{.Bio}},
        age: {{.Age}},
        inventory: {{index .Inventory 1}}
        friends {{.Friends}}
        address {{.Address}}
    }
    //console.log(user)
</script>