HOST := http://localhost:4566

#############
# DEBUG		#
#############
.PHONY: get-dynamo-tables
get-dynamo-tables:
	aws dynamodb list-tables --endpoint-url $(HOST)

#############
# TERRAFORM #
#############
.PHONY: run-tf
run-tf: 
	tflocal apply --auto-approve

