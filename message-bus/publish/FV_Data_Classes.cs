using System;
using System.Collections.Generic;
using System.Linq;
using System.Text;
using System.Threading.Tasks;

namespace CR_IO_fvStructures
{

    //detail members for use when publishing
    public struct PublishDetails
    {
        public string topicID { get; set; }
        public string topicName { get; set; }
        public string description { get; set; }
        public string flexField { get; set; }
    }

    //struct for patient data values
    public struct Patient
    {
        public string shortReport { get; set; }
        public bool intubated { get; set; }
        public int drips { get; set; }
        public int age { get; set; }
        public char gender { get; set; }
        public bool cardiac { get; set; }
        public bool giBleed { get; set; }
        public bool ob { get; set; }
    }

    //struct to allow for array of waypoints
    public struct Waypoint
    {
        public int id { get; set; }
        public string ete { get; set; }
        public string ett { get; set; }
        public bool active { get; set; }
    }

    //Primary class for mission
    public class Mission
    {
        public PublishDetails meta{ get; set; }
        public int missionID { get; set; }
        public string tcNUM { get; set; }
        public string asset { get; set; }
        public int reqID { get; set; }
        public int recID { get; set; }
        public string priority { get; set; }
        public string callType { get; set; }
        public Patient patient { get; set; }
        public int[] crewMemberIDs { get; set; }
        public Waypoint[] waypoints { get; set; }

        //Convert method to output JSON format as string using mission class instance
        //Uses a string builder to combine all values together with proper formatting
        //streamOUT initialy set to _FAIL to check that thread was built in powershell
        //-Value of _FAIL indicates to script that string did was not formed properly
        public string convertJSON()
        {
            string streamOUT = "_FAIL";
            StringBuilder strbWorker = new StringBuilder("");
            strbWorker.Append((string.Format("{{\"missionID\": \"{0}\", ", missionID)));
            strbWorker.Append((string.Format("\"TCNUM\": \"{0}\", ", tcNUM)));
            strbWorker.Append((string.Format("\"Asset\": \"{0}\", ", asset)));
            strbWorker.Append((string.Format("\"RegID\": \"{0}\", ", reqID)));
            strbWorker.Append((string.Format("\"RecID\": \"{0}\", ", recID)));
            strbWorker.Append((string.Format("\"Priority\": \"{0}\", ", priority)));
            strbWorker.Append((string.Format("\"CallType\": \"{0}\", ", callType)));
            strbWorker.Append((string.Format("\"Patient\": {{ ")));
            strbWorker.Append((string.Format("\"shortReport\": \"{0}\", ", patient.shortReport)));
            strbWorker.Append((string.Format("\"intubated\": \"{0}\", ", patient.intubated)));
            strbWorker.Append((string.Format("\"drips\": \"{0}\", ", patient.drips)));
            strbWorker.Append((string.Format("\"age\": \"{0}\", ", patient.age)));
            strbWorker.Append((string.Format("\"gender\": \"{0}\", ", patient.gender)));
            strbWorker.Append((string.Format("\"cardiac\": \"{0}\", ", patient.cardiac)));
            strbWorker.Append((string.Format("\"gibleed\": \"{0}\", ", patient.giBleed)));
            strbWorker.Append((string.Format("\"ob\": \"{0}\" ", patient.ob)));
            strbWorker.Append((string.Format(" }}, ")));
            try
            {
                //if to catch failure before writing crew member ID failure, otherwhise will not fail until loop
                if (crewMemberIDs.Length > 0)
                {
                    strbWorker.Append((string.Format("\"CrewMemberID\": [ ")));
                    //Loop through crewmember ids array, skip last for accounting line directly after
                    for (int i = 0; i < (crewMemberIDs.Length - 1); ++i)
                    {
                        strbWorker.Append((string.Format("\"{0}\", ", crewMemberIDs[i])));
                    }
                    //account for no comma at end of JSON arrays
                    strbWorker.Append((string.Format("\"{0}\" ", crewMemberIDs[(crewMemberIDs.Length - 1)])));
                    //end accounting
                    strbWorker.Append((string.Format("], ")));
                }
            }
            catch
            {
                Console.WriteLine("Error no crew defined for mission {0}", missionID);
            }
            try //allows processing to continue in case no waypoints defined
            {
                //if to catch failure before writing starting waypoints header
                if (waypoints.Length > 0)
                {
                    strbWorker.Append((string.Format("\"waypoints\": [ ")));
                    //loop through waypoints struct array, skip last element for account block directly after
                    for (int i = 0; i < (waypoints.Length - 1); ++i)
                    {
                        strbWorker.Append((string.Format("{{ ")));
                        strbWorker.Append((string.Format("\"ID\": \"{0}\", ", waypoints[i].id)));
                        strbWorker.Append((string.Format("\"ETE\": \"{0}\", ", waypoints[i].ete)));
                        strbWorker.Append((string.Format("\"ETT\": \"{0}\", ", waypoints[i].ett)));
                        strbWorker.Append((string.Format("\"Active\": \"{0}\" ", waypoints[i].active)));
                        strbWorker.Append((string.Format("}}, ")));
                    }
                    //Block to account for no comma at end of JSON arrays
                    strbWorker.Append((string.Format("{{ ")));
                    strbWorker.Append((string.Format("\"ID\": \"{0}\", ", waypoints[waypoints.Length - 1].id)));
                    strbWorker.Append((string.Format("\"ETE\": \"{0}\", ", waypoints[waypoints.Length - 1].ete)));
                    strbWorker.Append((string.Format("\"ETT\": \"{0}\", ", waypoints[waypoints.Length - 1].ett)));
                    strbWorker.Append((string.Format("\"Active\": \"{0}\" ", waypoints[waypoints.Length - 1].active)));
                    strbWorker.Append((string.Format("}} ")));
                    //end accounting block
                    strbWorker.Append((string.Format("], ")));
                }
            }
            catch
            {
                Console.WriteLine("Error no waypoints defined for mission {0}", missionID);
            }
            strbWorker.Append((string.Format("}}")));
            streamOUT = strbWorker.ToString();
            return streamOUT;
        }
    }

    public class Waypoints
    {
        public PublishDetails meta { get; set; }
        public int missionID { get; set; }
        Waypoint[] waypoints { get; set; }

        public string convertJSON()
        {
            string streamOUT = "_FAIL";
            StringBuilder strbWorker = new StringBuilder("");
            strbWorker.Append((string.Format("{{\"missionID\": \"{0}\", ", missionID)));
            try //allows processing to continue in case no waypoints defined
            {
                //if to catch failure before writing starting waypoints header
                if (waypoints.Length > 0)
                {
                    strbWorker.Append((string.Format("\"waypoints\": [ ")));
                    //loop through waypoints struct array, skip last element for account block directly after
                    for (int i = 0; i < (waypoints.Length - 1); ++i)
                    {
                        strbWorker.Append((string.Format("{{ ")));
                        strbWorker.Append((string.Format("\"ID\": \"{0}\", ", waypoints[i].id)));
                        strbWorker.Append((string.Format("\"ETE\": \"{0}\", ", waypoints[i].ete)));
                        strbWorker.Append((string.Format("\"ETT\": \"{0}\", ", waypoints[i].ett)));
                        strbWorker.Append((string.Format("\"Active\": \"{0}\" ", waypoints[i].active)));
                        strbWorker.Append((string.Format("}}, ")));
                    }
                    //Block to account for no comma at end of JSON arrays
                    strbWorker.Append((string.Format("{{ ")));
                    strbWorker.Append((string.Format("\"ID\": \"{0}\", ", waypoints[waypoints.Length - 1].id)));
                    strbWorker.Append((string.Format("\"ETE\": \"{0}\", ", waypoints[waypoints.Length - 1].ete)));
                    strbWorker.Append((string.Format("\"ETT\": \"{0}\", ", waypoints[waypoints.Length - 1].ett)));
                    strbWorker.Append((string.Format("\"Active\": \"{0}\" ", waypoints[waypoints.Length - 1].active)));
                    strbWorker.Append((string.Format("}} ")));
                    //end accounting block
                    strbWorker.Append((string.Format("], ")));
                }
            }
            catch
            {
                Console.WriteLine("Error no waypoints defined for mission {0}", missionID);
            }
            strbWorker.Append((string.Format("}}")));
            streamOUT = strbWorker.ToString();
            return streamOUT;
        }
    }

    public class MissionCrewUpdate
    {
        public PublishDetails meta { get; set; }
        public string missionID { get; set; }
        public int[] crewMemberIDs { get; set; }

        public string convertJSON()
        {
            string streamOUT = "_FAIL";
            StringBuilder strbWorker = new StringBuilder("");
            strbWorker.Append((string.Format("{{\"missionID\": \"{0}\", ", missionID)));
            try
            {
                //if to catch failure before writing crew member ID failure, otherwhise will not fail until loop
                if (crewMemberIDs.Length > 0)
                {
                    strbWorker.Append((string.Format("\"CrewMemberID\": [ ")));
                    //Loop through crewmember ids array, skip last for accounting line directly after
                    for (int i = 0; i < (crewMemberIDs.Length - 1); ++i)
                    {
                        strbWorker.Append((string.Format("\"{0}\", ", crewMemberIDs[i])));
                    }
                    //account for no comma at end of JSON arrays
                    strbWorker.Append((string.Format("\"{0}\" ", crewMemberIDs[(crewMemberIDs.Length - 1)])));
                    //end accounting
                    strbWorker.Append((string.Format("], ")));
                }
            }
            catch
            {
                Console.WriteLine("Error no crew defined for mission {0}", missionID);
            }
            strbWorker.Append((string.Format("}}")));
            streamOUT = strbWorker.ToString();
            return streamOUT;
        }
    }

    public class WaypointDetail
    {
        public PublishDetails meta { get; set; }
        public int ID { get; set; }
        public string notes { get; set; }
        public string name { get; set; }
        public string type { get; set; }
        public string address1 { get; set; }
        public string address2 { get; set; }
        public string country { get; set; }
        public string state { get; set; }
        public string county { get; set; }
        public string city { get; set; }
        public string zip { get; set; }
        public Double gpslat { get; set; }
        public Double gpslong { get; set; }
        public int gpsWaypoint { get; set; }
        public string airportIdentifier { get; set; }
        public string[] phones { get; set; }
        public string shortCode { get; set; }
        public string padTime { get; set; }
        public string[] radioChannels { get; set; }
        public string notAMS { get; set; }

        public string convertJSON()
        {
            string streamOUT = "_FAIL";
            StringBuilder strbWorker = new StringBuilder("");
            strbWorker.Append((string.Format("{{\"ID\": \"{0}\", ", ID)));
            strbWorker.Append((string.Format("\"notes\": \"{0}\", ", notes)));
            strbWorker.Append((string.Format("\"name\": \"{0}\", ", name)));
            strbWorker.Append((string.Format("\"type\": \"{0}\", ", type)));
            strbWorker.Append((string.Format("\"address1\": \"{0}\", ", address1)));
            strbWorker.Append((string.Format("\"address2\": \"{0}\", ", address2)));
            strbWorker.Append((string.Format("\"country\": \"{0}\", ", country)));
            strbWorker.Append((string.Format("\"state\": \"{0}\", ", state)));
            strbWorker.Append((string.Format("\"county\": \"{0}\", ", county)));
            strbWorker.Append((string.Format("\"city\": \"{0}\", ", city)));
            strbWorker.Append((string.Format("\"zip\": \"{0}\", ", zip)));
            strbWorker.Append((string.Format("\"lat\": \"{0}\", ", gpslat)));
            strbWorker.Append((string.Format("\"long\": \"{0}\", ", gpslong)));
            strbWorker.Append((string.Format("\"GPSWaypoint\": \"{0}\", ", gpsWaypoint)));
            strbWorker.Append((string.Format("\"AirportIdentifier\": \"{0}\", ", airportIdentifier)));
            try
            {
                //if to catch failure before writing phone array, otherwhise will not fail until loop
                if (phones.Length > 0)
                {
                    strbWorker.Append((string.Format("\"phone\": [ ")));
                    //Loop through phones array, skip last for accounting line directly after
                    for (int i = 0; i < (phones.Length - 1); ++i)
                    {
                        strbWorker.Append((string.Format("\"{0}\", ", phones[i])));
                    }
                    //account for no comma at end of JSON arrays
                    strbWorker.Append((string.Format("\"{0}\" ", phones[(phones.Length - 1)])));
                    //end accounting
                    strbWorker.Append((string.Format("], ")));
                }
            }
            catch
            {
                Console.WriteLine("Error no phones for waypoint {0}", ID);
            }
            try //allows processing to continue in case no radios defined
            {
                //if to catch failure before writing starting radios header
                if (radioChannels.Length > 0)
                {
                    strbWorker.Append((string.Format("\"radioChannels\": [ ")));
                    //Loop through radios array, skip last for accounting line directly after
                    for (int i = 0; i < (radioChannels.Length - 1); ++i)
                    {
                        strbWorker.Append((string.Format("\"{0}\", ", radioChannels[i])));
                    }
                    //account for no comma at end of JSON arrays
                    strbWorker.Append((string.Format("\"{0}\" ", radioChannels[(radioChannels.Length - 1)])));
                    //end accounting
                    strbWorker.Append((string.Format("], ")));
                }
            }
            catch
            {
                Console.WriteLine("Error no radios defined for waypoint {0}", ID);
            }
            strbWorker.Append((string.Format("}}")));
            streamOUT = strbWorker.ToString();
            return streamOUT;

        }
    }

    public class Aircraft
    {
        public PublishDetails meta { get; set; }
        public int ID { get; set; }
        public string tailNum { get; set; }
        public string satPhone { get; set; }
        public string cellPhone { get; set; }
        public string baseID { get; set; }
        public string callsign { get; set; }
        public int maxPatientWieght { get; set; }
        public int padTimeDay { get; set; }
        public int padTimeNight { get; set; }
        public string vendor { get; set; }
        public string status { get; set; }
        public string specialEquipment { get; set; }
        public string color { get; set; }
        public int lastPOS { get; set; }
        public string model { get; set; }
        public int[] callTypes { get; set; }

        public string convertJSON()
        {
            string streamOUT = "_FAIL";
            StringBuilder strbWorker = new StringBuilder("");
            strbWorker.Append((string.Format("{{\"ID\": \"{0}\", ", ID)));
            strbWorker.Append((string.Format("\"nNum\": \"{0}\", ", tailNum)));
            strbWorker.Append((string.Format("\"satPhone\": \"{0}\", ", satPhone)));
            strbWorker.Append((string.Format("\"cellphone\": \"{0}\", ", cellPhone)));
            strbWorker.Append((string.Format("\"baseID\": \"{0}\", ", baseID)));
            strbWorker.Append((string.Format("\"callsign\": \"{0}\", ", callsign)));
            strbWorker.Append((string.Format("\"maxPatientWeight\": \"{0}\", ", maxPatientWieght)));
            strbWorker.Append((string.Format("\"padTimeDay\": \"{0}\", ", padTimeDay)));
            strbWorker.Append((string.Format("\"padTimeNight\": \"{0}\", ", padTimeNight)));
            strbWorker.Append((string.Format("\"vendor\": \"{0}\", ", vendor)));
            strbWorker.Append((string.Format("\"status\": \"{0}\", ", status)));
            strbWorker.Append((string.Format("\"specialEquipment\": \"{0}\", ", specialEquipment)));
            strbWorker.Append((string.Format("\"color\": \"{0}\", ", color)));
            strbWorker.Append((string.Format("\"lastKnownLocation\": \"{0}\", ", lastPOS)));
            strbWorker.Append((string.Format("\"model\": \"{0}\", ", model)));
            try
            {
                //if to catch failure before writing calltypes array, otherwhise will not fail until loop
                if (callTypes.Length > 0)
                {
                    strbWorker.Append((string.Format("\"callTypes\": [ ")));
                    //Loop through phones array, skip last for accounting line directly after
                    for (int i = 0; i < (callTypes.Length - 1); ++i)
                    {
                        strbWorker.Append((string.Format("\"{0}\", ", callTypes[i])));
                    }
                    //account for no comma at end of JSON arrays
                    strbWorker.Append((string.Format("\"{0}\" ", callTypes[(callTypes.Length - 1)])));
                    //end accounting
                    strbWorker.Append((string.Format("], ")));
                }
            }
            catch
            {
                Console.WriteLine("Error no calltypes for aircraft {0}", ID);
            }
            strbWorker.Append((string.Format("}}")));
            streamOUT = strbWorker.ToString();
            return streamOUT;
        }
    }
}
