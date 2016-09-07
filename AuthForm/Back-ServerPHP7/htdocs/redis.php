<?php
echo "Loading redis...<br>";
$services = json_decode($_ENV["VCAP_SERVICES"]);
//Use iteration on services instead
try {
    $redis = new Redis();
    $redis->connect('10.234.254.96', '6379');
    $redis->auth('red!s');
} catch (Exception $e) {
    die('Error : ' . $e->getMessage());
}
echo "<br>Connected to Redis...<br>" . var_dump($services); 
// we create the table if not exists
if (!empty($_GET["key"]) && !empty($_GET["value"])) {
    $redis->set($_GET["key"], $_GET["value"]);
}
?>
