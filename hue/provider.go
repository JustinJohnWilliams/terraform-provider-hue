package hue

import (
	"reflect"
	"regexp"
	"strings"

	"github.com/hashicorp/terraform-plugin-sdk/helper/schema"
	"github.com/hashicorp/terraform-plugin-sdk/terraform"
)

func Provider() terraform.ResourceProvider {
	p := &schema.Provider{
		Schema: map[string]*schema.Schema{
			"host": {
				Type:        schema.TypeString,
				Required:    true,
				Description: "Domain name/IP of the Hue Bridge",
				DefaultFunc: schema.EnvDefaultFunc("HUE_HOST", nil),
			},
			"username": {
				Type:        schema.TypeString,
				Required:    true,
				Description: "Username of Hue Bridge",
				DefaultFunc: schema.EnvDefaultFunc("HUE_USER", nil),
			},
		},
		DataSourcesMap: map[string]*schema.Resource{
			"hue_lights": dataSourceHueLights(),
			"hue_light":  dataSourceHueLight(),
		},
		ResourcesMap: map[string]*schema.Resource{
			"hue_light": resourceHueLight(),
		},
	}
	p.ConfigureFunc = func(d *schema.ResourceData) (interface{}, error) {
		terraformVersion := p.TerraformVersion
		if terraformVersion == "" {
			// Terraform 0.12 introduced this field to the protocol
			// We can therefore assume that if it's missing it's 0.10 or 0.11
			terraformVersion = "0.11+compatible"
		}
		return providerConfigure(d, terraformVersion)
	}
	return p
}

func providerConfigure(d *schema.ResourceData, terraformVersion string) (interface{}, error) {
	config := Config{
		Host:     d.Get("host").(string),
		Username: d.Get("username").(string),
	}

	cfg, err := config.Client()
	if err != nil {
		return cfg, err
	}

	return cfg, err
}

//Convert slice of strings to schema.TypeSet
func makeStringList(list *[]string) []interface{} {
	ilist := make([]interface{}, len(*list))
	for i, v := range *list {
		ilist[i] = v
	}
	return ilist
}

//Convert slice of strings to schema.Set
func makeStringSet(list *[]string) *schema.Set {
	ilist := make([]interface{}, len(*list))
	for i, v := range *list {
		ilist[i] = v
	}
	return schema.NewSet(schema.HashString, ilist)
}

//Convert schema.TypeList to a slice of strings
func listToStringSlice(s []interface{}) []string {
	list := make([]string, len(s))
	for i, v := range s {
		list[i] = v.(string)
	}
	return list
}

//Convert schema.Set to a slice of strings
func setToStringSlice(s *schema.Set) []string {
	list := make([]string, s.Len())
	for i, v := range s.List() {
		list[i] = v.(string)
	}
	return list
}

//Copy map values into an object where map key == object field name (e.g. map[foo] == &{Foo: ...}
func mapEntity(d map[string]interface{}, obj interface{}) {
	val := reflect.ValueOf(obj).Elem()
	for field := range d {
		f := val.FieldByName(strings.Title(field))
		if f.IsValid() {
			if f.Kind() == reflect.Slice {
				incoming := d[field].([]interface{})
				s := reflect.MakeSlice(f.Type(), len(incoming), len(incoming))
				for i := 0; i < len(incoming); i++ {
					s.Index(i).Set(reflect.ValueOf(incoming[i]))
				}
				f.Set(s)
			} else {
				f.Set(reflect.ValueOf(d[field]))
			}
		} else {
			f := val.FieldByName(strings.Title(toCamelCase(field)))
			f.Set(reflect.ValueOf(d[field]))
		}
	}
}

//Break a string in the format /Partition/name into a Partition / Name object
func parseF5Identifier(str string) (partition, name string) {
	if strings.HasPrefix(str, "/") {
		ary := strings.SplitN(strings.TrimPrefix(str, "/"), "/", 2)
		return ary[0], ary[1]
	}
	return "", str
}

// Convert Snakecase to Camelcase
func toCamelCase(str string) string {
	var link = regexp.MustCompile("(^[A-Za-z])|_([A-Za-z])")
	return link.ReplaceAllStringFunc(str, func(s string) string {
		return strings.ToUpper(strings.Replace(s, "_", "", -1))
	})
}

// Convert Camelcase to Snakecase
func toSnakeCase(str string) string {
	var matchFirstCap = regexp.MustCompile("(.)([A-Z][a-z]+)")
	var matchAllCap = regexp.MustCompile("([a-z0-9])([A-Z])")
	snake := matchFirstCap.ReplaceAllString(str, "${1}_${2}")
	snake = matchAllCap.ReplaceAllString(snake, "${1}_${2}")
	return strings.ToLower(snake)
}
