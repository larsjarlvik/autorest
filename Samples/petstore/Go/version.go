package petstore


import (
    "fmt"
)

const (
    major = "0"
    minor = "0"
    patch = "0"
    // Always begin a "tag" with a dash (as per http://semver.org)
    tag   = "-beta"
    semVerFormat = "%s.%s.%s%s"
    userAgentFormat = "Azure-SDK-For-Go/%s arm-%s/%s"
)

// UserAgent returns the UserAgent string to use when sending http.Requests.
func UserAgent() string {
    return fmt.Sprintf(userAgentFormat, Version(), "petstore", "1.0.0")
}

// Version returns the semantic version (see http://semver.org) of the client.
func Version() string {
    return fmt.Sprintf(semVerFormat, major, minor, patch, tag)
}
