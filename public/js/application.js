AWS.config.update({region: 'us-west-2'});
AWS.config.credentials = new AWS.CognitoIdentityCredentials({IdentityPoolId: 'us-west-2:d49682ff-eebe-4370-aeea-a37115092b61'});

var lambdaParams = { region: 'us-west-2', apiVersion: '2015-03-31' };
var lambda = new AWS.Lambda(lambdaParams);
