using Google.Protobuf;
using Google.Cloud.PubSub.V1;
using Google.Apis.Auth;
using Grpc.Core;
using System;
using System.Collections.Generic;
using System.Linq;
using System.Text;
using System.Threading.Tasks;

namespace CR_IO_GGL_PubSub_DB
{
    class Program
    {
        static void Main(string[] args)
        {
            string callProjID = args[0];
            string callTopic = args[1];
            string callData = args[2];
            string callAttributeKey = args[3];
            string callAttributeVal = args[4];
            //Check caller variables
            if (callProjID == null || callProjID == "")
            {
                throw new System.ArgumentException("Parameter is NULL", "callProjID");
            }

            if (callTopic == null || callTopic == "")
            {
                throw new System.ArgumentException("Parameter is NULL", "callTopic");
            }

            if (callData == null || callData == "")
            {
                throw new System.ArgumentException("Parameter is NULL", "callData");
            }

            if (callAttributeKey == null || callAttributeKey == "")
            {
                throw new System.ArgumentException("Parameter is NULL", "callAttributeKey");
            }

            if (callAttributeVal == null || callAttributeVal == "")
            {
                throw new System.ArgumentException("Parameter is NULL", "callAttributeVal");
            }

            //Create topic name
            var topicName = new TopicName(callProjID, callTopic);

            //Create PublisherAPIClient
            PublisherServiceApiClient localPublisher = PublisherServiceApiClient.Create();

            //Create Message
            var localMessage = new PubsubMessage()
            {
                Data = ByteString.FromBase64(callData)
            };

            localMessage.Attributes.Add(callAttributeKey, callAttributeVal);

            //Create Message List
            var localMsgList = new List<PubsubMessage> { localMessage };

            //Publish Message
            localPublisher.Publish(topicName, localMsgList);
        }
    }
}
