# SmallWatch
SmallWatch is a lightweight Amazon Web Services (AWS) CloudWatch management tool written in Go that provides list/reduce/delete capabilities within a single account.

## Features
* __list__: Lists out all CloudWatch Log Groups within the account.
* __reduce__: Allows modification of CloudWatch Log Group retention times.
* __delete__: Allows deletion of N-number log groups based on prefix.

## Installation
*  `go install github.com/hunttom/smallWatch`

## QuickStart
```bash
$smallWatch

smallwatch v0.0.1 - Cloudwatch log cleanup for Dev Environments

Available commands:

   list     List out CloudWatch logs 
   reduce   Reduce CloudWatch logs retention size 
   delete   List of CloudWatch Logs to delete 

Flags:

  -help
    	Get help on the 'smallwatch' command.
```

## Usage Examples
### List:

Help command
```bash
$smallWatch list --help
smallwatch v0.0.1 - Cloudwatch log cleanup for Dev Environments

smallwatch list - List out CloudWatch logs
Flags:

  -help
        Get help on the 'smallwatch list' command.
  -prefix string
        The prefix of the log group to list (default "/")
```

Run command
```bash
$smallWatch list
Log Group: /aws/codebuild/foo
Log Group: /aws/codebuild/foo-b
Log Group: /aws/codebuild/bar
Log Group: /aws/codebuild/biz
Found 4 log groups
```

### Reduce:
Help command
```bash
$smallWatch reduce --help
smallwatch v0.0.1 - Cloudwatch log cleanup for Dev Environments

smallwatch reduce - Reduce CloudWatch logs retention size
Flags:

  -days int
        The number of days to set retention. Defaults to 30 days. (default 30)
  -help
        Get help on the 'smallwatch reduce' command.
  -no-dry-run
        Performs actual modification
  -prefix string
        The prefix of the log group to reduce
```

Dry Run command
```bash
$smallWatch reduce -days 7 -prefix "/"
DRY RUN - Reducing Log Group: /aws/codebuild/foo
DRY RUN - Reducing Log Group: /aws/codebuild/foo-b
DRY RUN - Reducing Log Group: /aws/codebuild/bar
DRY RUN - Reducing Log Group: /aws/codebuild/biz
Set 4 log groups to 7 days retention
```

No-Dry Run command
```bash
$smallWatch reduce -days 7 -prefix "/" -no-dry-run
Reducing Log Group: /aws/codebuild/foo
Reducing Log Group: /aws/codebuild/foo-b
Reducing Log Group: /aws/codebuild/bar
Reducing Log Group: /aws/codebuild/biz
Set 4 log groups to 7 days retention
```

### Delete:
Help command
```bash
$smallWatch delete --help
smallwatch v0.0.1 - Cloudwatch log cleanup for Dev Environments

smallwatch delete - List of CloudWatch Logs to delete
Flags:

  -help
        Get help on the 'smallwatch delete' command.
  -no-dry-run
        Performs actual deletion
  -prefix string
        The prefix of the log group to delete
```

Dry Run command
```bash
$smallWatch delete -prefix /aws/codebuild/foo
DRY RUN - Deleting Log Group: /aws/codebuild/foo
DRY RUN - Deleting Log Group: /aws/codebuild/foo-b
Deleted 2 log groups
```

No-Dry Run command
```
$smallWatch delete -prefix /aws/codebuild/foo -no-dry-run
Deleting Log Group: /aws/codebuild/foo
Deleting Log Group: /aws/codebuild/foo-b
Deleted 2 log groups
```

## Contributing
* Fork the repository
* Create your feature branch ( git checkout -b feature/amazing-feature)
* Commit your changes ( git commit -m 'Add some amazing feature')
* Push to the branch ( git push origin feature/amazing-feature)
* Open a Pull Request

## License
This project is licensed under the MIT License - see the LICENSE file for details.

## Acknowledgments
Built with Go

Uses CLIR as a CLI framework: https://github.com/leaanthony/clir

## Support
For support, please open an issue in the GitHub repository or contact the maintainers.