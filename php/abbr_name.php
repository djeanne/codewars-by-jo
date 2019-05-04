<?php

function abbrevName($name) {
	$abbr = explode(" ", $name);
	$finit = strtoupper($abbr[0][0]);
	$linit = strtoupper($abbr[1][0]);
	return "{$finit}.{linit}"
}


