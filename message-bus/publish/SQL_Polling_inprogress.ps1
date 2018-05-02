#CR SQL POLLING/PUBLISH - BENCHMARKS INITIAL (POLLING MACHINE - 2% CPU 0% NETWORK) -> (SQL MACHINE - 1% CPU 0% NETWORK) --> ON RYZEN 16 Core 3.8GHz
#TODO - SQL QUERIES FOR TOPICS
#ADD LOGIC CODE FOR GOOGLE PUB/SUB
#ADD RUNSPACE APARTMENT LOGIC
#ADD APPLICATION CONTEXT IF NEEDED

$ref_SQL_TRANSACTION_LOG = "
SELECT 
[Operation],
[AllocUnitName],
[Slot ID]
FROM sys.fn_dblog(NULL,NULL)
WHERE Operation IN ('LOP_INSERT_ROWS','LOP_MODIFY_ROW',
    'LOP_DELETE_ROWS')
AND AllocUnitName IN ('dbo.Missions.pk_missions')
"

function fProcess_Change($callOperation, $callUnitName, $callSlotID) {
    Write-Host "OPERATION-$callOperation IN TABLE-$callUnitName IN ROW-$callSlotID"
}

[int]$Last_Length = 0
$g_Error_Level = 0

$ref_SQL_SERVER = "CR-DEV-SQL"
$ref_SQL_INSTANCE = "DEVSQL"
$ref_SQL_DATABASE = "FlightVector"
$ref_Poll_Interval = 1000 #in milliseconds

$ref_Connection_String = "Server=$ref_SQL_SERVER\$ref_SQL_INSTANCE;Database=$ref_SQL_DATABASE;Integrated Security=True"

[int]$p_int = 0

$TEST_SQL_Connection = New-Object System.Data.SqlClient.SqlConnection($ref_Connection_String)
$TEST_SQL_Connection.Open()
$TEST_SQL_COMMAND = New-Object System.Data.SqlClient.SqlCommand($ref_SQL_TRANSACTION_LOG,$TEST_SQL_Connection)
$TEST_SQL_Adapter = New-Object System.Data.SqlClient.SqlDataAdapter $TEST_SQL_COMMAND
$TEST_DATA_SET = New-Object System.Data.DataSet
While($g_Error_Level -ge 0) {
    $TEST_DATA_SET = New-Object System.Data.DataSet
    $TEST_SQL_Adapter.Fill($TEST_DATA_SET) | Out-Null
    if($Last_Length -ne $TEST_DATA_SET.Tables.Operation.Length) {
        if($Last_Length -gt $TEST_DATA_SET.Tables.Operation.Length) {
            #PROCESS ALL
            $p_int = 0
        }
        elseif($Last_Length -lt $TEST_DATA_SET.Tables.Operation.Length) {
            #PROCESS DIFFERENCE
            $p_int = $Last_Length
        }
        While($p_int -lt $TEST_DATA_SET.Tables.Operation.Length) {
            fProcess_Change $TEST_DATA_SET.Tables.Operation[$p_int] $TEST_DATA_SET.Tables.AllocUnitName[$p_int] $TEST_DATA_SET.Tables.SlotID[$p_int]
            $p_int ++
        }
    }
    $Last_Length = $TEST_DATA_SET.Tables.Operation.Length
    # SKIP IF EQUAL
    [System.GC]::Collect()
    Start-Sleep -Milliseconds $ref_Poll_Interval
}