<?php

$sessionId = $_POST["sessionId"];
$serviceCode =$_POST["serviceCode"];
$phoneNumber = $_POST["phoneNumber"];
$text = $_POST["text"];

if ($text =="") {
	//This is the first request
	$response = "CON What would you like to check? \n";
	$response .= "1. Covid-19 symptoms \n";
	$response .= "2. How to handle a covid-19 patient \n";
	$response .= "3. Stress Management \n";	
	$response .= "4. Other";
}else if($text == "1"){
	$response = "END The following are the Covid-19 symptoms \n
	 1.Fever \n
	 2.Dry cough \n
	 3.Tiredness \n";
	
}else if($text == "2"){
	$response = "END This is how you handle a Covid-19 patient \n
	 1.Please visit the following webpage for more information https://www.who.int/emergencies/diseases/novel-coronavirus-2019/technical-guidance/patient-management \n";
	
}else if($text == "3"){
	$response = "CON Erevusha is here to help out on stress and trauma management among the front-line health care workers, pick an option: \n";
	$response .= "1. Talk to our chatbot \n"; 
	$response .= "2.Join a whatsapp group \n";
	
}else if($text == "3*2"){
	//$docs = https://chat.whatsapp.com/BmgK2L1dKfv2DcnIGK287u 
	//$nurses =https://chat.whatsapp.com/CdhBpIgMuxS9rtkykUox27
	//$CHW = https://chat.whatsapp.com/JtZ9rv58ACUKc7tsDMwiTT
	$response = "CON What is your occupation? \n";
	$response .= "1.Doctor \n";
	$response .= "2.Nurse \n";
	$response .= "3.CHW \n";
}else if($text == "3*2*1"){
	$response = "END Click to join this whatsapp group \n
	 https://chat.whatsapp.com/BmgK2L1dKfv2DcnIGK287u";
}else if($text == "3*2*2"){
	$response = "END Click to join this whatsapp group \n
	 https://chat.whatsapp.com/CdhBpIgMuxS9rtkykUox27";
}else if($text == "3*2*3"){
	$response = "END Click to join this whatsapp group \n
	 https://chat.whatsapp.com/JtZ9rv58ACUKc7tsDMwiTT";
}else if($text == "4"){
	$response = "END Kindly dial the following to talk directly with the MOH \n
	 *719# komeshaCorona";
}

header('Content-type: text/plain');
echo $response;
	

?>
