module thomasbeukema/gibberish

go 1.17

require gopkg.in/yaml.v2 v2.4.0

require internal/config v1.0.0

replace internal/config => ./internal/conf

require internal/api v1.0.0

replace internal/api => ./internal/api
