# slack-emoji-reminder

## Deploy

```sh
# create Cloud Pub/Sub
gcloud pubsub topics create topic-slack-emoji-reminder --project <YOUR GCP PROJECT>

# create Cloud Functions
gcloud functions deploy SlackEmojiReminder --project <YOUR GCP PROJECT> \
  --entry-point Remind \
  --trigger-topic topic-slack-emoji-reminder \
  --runtime go113 \
  --set-env-vars REACTION=<YOUR TARGET REACTION,EMOJI> \
  --set-env-vars USER=<YOUR USER ID> \
  --set-env-vars SLACK_BOT_TOKEN=<YOUR BOT TOKEN WITH WRITE AUTH> \
  --set-env-vars SLACK_REMIND_CHANNEL=<YOUR TARGET REMIND CHANNEL> \
  --set-env-vars SLACK_API_TOKEN=<YOUR API TOKEN WITH SEARCH AUTH> \
  --set-env-vars SLACK_SEARCH_CHANNEL=<YOUR TARGET SEARCH CHANNEL>

# create Cloud Scheduler
gcloud beta scheduler jobs create pubsub slack-emoji-reminder --project <YOUR GCP PROJECT> \
  --schedule "0 8 * * *" \
  --topic topic-slack-emoji-reminder \
  --message-body="execute" \
  --time-zone "Asia/Tokyo"
```
