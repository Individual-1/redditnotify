# Reddit Notifier

## Overview
This is a basic notification app that polls Reddit for new posts from the subreddits and users you specify.

## Architecture

### Polling Service
The core of this design is the polling service which reads a configuration file containing a list of subreddits and users to follow, then either pulls each individually on a timer, or utilizes an account's friends and home functionality to pull new items in fewer requests.

#### Configuration

##### Specifying Subreddits
Json array

#### Specifying Users
Json array of objects with each User object allowing for whitelist or blacklist approach for subreddits to include (boolean indicating type)

#### Polling Logic

For both methods, poll the desired endpoints and cache last seen ids. If filtering user subreddits then apply the appropriate filter after fetching the user posts and comments.

##### No user account
Loop over endpoints for individual subreddits and user profile posts/comments.

##### User account provided
On launch, load the list of subreddits and users then verify that the account provided has friended those users or followed those accounts. From there, track the latest sub posts via `reddit.com/new/` and friend posts/comments via `reddit.com/r/friends` and `reddit.com/r/friends/comments`
