package generated

import (
	"errors"
	yaml "gopkg.in/yaml.v2"
	appsv1 "k8s.io/api/apps/v1"
	corev1 "k8s.io/api/core/v1"
	"k8s.io/apimachinery/pkg/runtime"
	"k8s.io/client-go/kubernetes/scheme"
)

// SpinnakerGeneratedConfig represents the config generated by Halyard
type SpinnakerGeneratedConfig struct {
	Config map[string]ServiceConfig `json:"config"`
}

// ServiceConfig is the generated service config
type ServiceConfig struct {
	Deployment *appsv1.Deployment `json:"deployment,omitempty"`
	Service    *corev1.Service    `json:"service,omitempty"`
	Resources  []runtime.Object   `json:"resources,omitempty"`
	ToDelete   []runtime.Object   `json:"todelete,omitempty"`
}

func (r *ServiceConfig) UnmarshalYAML(unmarshal func(interface{}) error) error {
	a := make(map[string]interface{})
	err := unmarshal(&a)

	if err != nil {
		return err
	}
	dser := scheme.Codecs.UniversalDeserializer()
	val := a["deployment"]
	if val != nil {
		o, err := translate(val, dser)
		if err != nil {
			return err
		}
		d, ok := o.(*appsv1.Deployment)
		if !ok {
			return errors.New("Invalid deployment")
		}
		r.Deployment = d
	}

	val = a["service"]
	if val != nil {
		o, err := translate(val, dser)
		if err != nil {
			return err
		}
		s, ok := o.(*corev1.Service)
		if !ok {
			return errors.New("Invalid service")
		}
		r.Service = s
	}

	val = a["resources"]
	l, ok := val.([]interface{})
	if ok {
		rs := make([]runtime.Object, 0)
		for i := range l {
			o, err := translate(l[i], dser)
			if err != nil {
				return err
			}
			rs = append(rs, o)
		}
		r.Resources = rs
	}
	return nil
}

func translate(decoded interface{}, decode runtime.Decoder) (runtime.Object, error) {
	b, err := yaml.Marshal(decoded)
	if err != nil {
		return nil, err
	}
	obj, _, err := decode.Decode(b, nil, nil)
	return obj, err
}
