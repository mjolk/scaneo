package main

const (
	scansText = `{{define "scans"}}// DON'T EDIT *** generated by scaneo *** DON'T EDIT //

package {{.PackageName}}

import "database/sql"

{{range .Tokens}}
{{ if .Composite}}
type proxy{{.Name}} struct {
	*{{.Name}}{{range .ProxyTypes}}
	proxy{{.}} *{{.}}{{end}}
}

func {{$.Visibility}}et{{title .Name}}s(rs *sql.Rows) ([]*{{.Name}}, error) {
	var err error
	result := make([]*{{.Name}}, 25)
	for rs.Next() {
		exists := false
		s := proxy{{.Name}}{ {{.Name}}:&{{.Name}}{}, {{range .ProxyTypes}}
		proxy{{.}}:&{{.}}{},{{end}}
		}
		if err = rs.Scan({{range .Fields}}
			&s.{{.Name}},{{end}}
		); err != nil {
			return nil, err
		}
		entity := s.{{.Name}}
		var newID bool = true
		for _, ent := range result {
			if ent.ID == s.ID {
				entity = ent
				newID = false
			}
		}
		{{$fields:=.Fields}}{{range .ProxyTypes}}exists = false
		if s.proxy{{.}}.ID != 0 {
			if len(entity.{{.}}) > 0 {
				for _, rel := range entity.{{.}} {
					if rel.ID == s.proxy{{.}}.ID {
						exists = true
					}
				}
			}
		}
		if !exists {
			{{$i:=.}}new{{.}} := &{{.}}{
				{{range $fields}}{{$prox:=proxy .Name}}{{if eq $prox $i}}{{field .Name}}: s.proxy{{$i}}.{{field .Name}},{{end}}{{end}}
			}
			entity.{{.}} = append(entity.{{.}}, new{{.}})
		}
		{{end}}if newID {
			result = append(result, entity)
		}
	}
	if err = rs.Err(); err != nil {
		return nil, err
	}
	return result, nil
}
{{else}}
func {{$.Visibility}}et{{title .Name}}(r *sql.Row) ({{.Name}}, error) {
	var s {{.Name}}
	if err := r.Scan({{range .Fields}}
		&s.{{.Name}},{{end}}
	); err != nil {
		return {{.Name}}{}, err
	}
	return s, nil
}

func {{$.Visibility}}et{{title .Name}}s(rs *sql.Rows) ([]*{{.Name}}, error) {
	structs := make([]*{{.Name}}, 0, 16)
	var err error
	for rs.Next() {
		var s {{.Name}}
		if err = rs.Scan({{range .Fields}}
			&s.{{.Name}},{{end}}
		); err != nil {
			return nil, err
		}
		structs = append(structs, &s)
	}
	if err = rs.Err(); err != nil {
		return nil, err
	}
	return structs, nil
}

{{end}}{{end}}
{{end}}`
)
