workflow "New workflow" {
  on = "push"
  resolves = ["docker://golang:1.12rc1"]
}

action "docker://golang:1.12rc1" {
  uses = "docker://golang:1.12rc1"
  args = "go build ./..."
}
