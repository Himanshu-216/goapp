{{- define "goapp.name" -}}
goapp
{{- end -}}

{{- define "goapp.fullname" -}}
{{ .Release.Name }}-{{ include "goapp.name" . }}
{{- end -}}
