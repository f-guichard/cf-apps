<?php
try {
    $redis = new Redis();
    $redis->connect($redisInfo["host"], $redisInfo["port"]);
    $redis->auth($redisInfo["password"]);
} catch (Exception $e) {
    die('Error : ' . $e->getMessage());
}
// we create the table if not exists
if (!empty($_POST["key"]) && !empty($_POST["value"])) {
    $redis->set($_POST["key"], $_POST["value"]);
}
$keys = $redis->keys("*");
?>
