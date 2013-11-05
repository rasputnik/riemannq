Dead simple query tool for [Riemann](http://riemann.io). 

Outputs JSON for easy slicing and dicing with tools like [jq](http://stedolan.github.io/jq/).

Usage
-----

    riemannq -s your.riemann.server:5555 -q 'tagged = "cpu"'

    riemannq -q 'tagged = "io"' | jq '.[] | {Host, Service, Metric}'