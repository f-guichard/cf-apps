<?php
  //Sample php7 code
  declare(strict_types=1);

  //Throw TypeExcpetion if (! castVar typeof string)
  function php7strictString(string $castVar) : string
  {
    return $castVar . " is a string";
  }

  try 
  {
    echo php7strictString("CFY");
    echo php7strictString(7);
    echo unknownfunction();
  }
  catch (EngineException exp)
  {
    echo var_dump($exp);
  }
  catch (TypeException exp)
  {
    echo var_dump($exp);
  }

  $vcapServices = json_decode($_ENV['VCAP_SERVICES'], true);
  $services = array();
  if(!isset($vcapServices))
  {
    echo "{ vcap_service : null}";
  }
  else
  {
    echo var_dump($vcapServices);
  }
?>
