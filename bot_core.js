var bk = require('botkit');
var controller = bk.slackbot();
var bot = controller.spawn({
	token:xoxb-18722935172-R5J4OCvGUgkY02Mhfphb0Kd3
});
bot.startRTM(function(err,bot,payload){
	if(err){
		throw new Error('Could not connect to slack');
	}
	console.log("here");
});
console.log("here2");