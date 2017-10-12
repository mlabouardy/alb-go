resource "aws_alb" "default" {
  name = "alb-api"
  internal = false
  security_groups = ["${aws_security_group.alb_sg.id}"]

}
