# slack-emoji-reminder

## Environment

```sh
export REACTION=<YOUR TARGET REACTION,EMOJI>
export USER=<YOUR USER ID>
export SLACK_BOT_TOKEN=<YOUR BOT TOKEN WITH WRITE AUTH>
export SLACK_REMIND_CHANNEL=<YOUR TARGET REMIND CHANNEL>
export SLACK_API_TOKEN=<YOUR API TOKEN WITH SEARCH AUTH>
export SLACK_SEARCH_CHANNEL=<YOUR TARGET SEARCH CHANNEL>
```

## Deploy

```sh
# create Cloud Pub/Sub
gcloud pubsub topics create topic-slack-emoji-reminder --project <YOUR GCP PROJECT>

# create Cloud Functions
gcloud functions deploy SlackEmojiReminder --project <YOUR GCP PROJECT> \
  --entry-point Remind \
  --trigger-topic topic-slack-emoji-reminder \
  --runtime go113
```