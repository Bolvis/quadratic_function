<?php
$a = (integer)$_POST['a'];
$b = (integer)$_POST['b'];
$c = (integer)$_POST['c'];

if($a == 0){
    echo "błędne pole A";
}else{
    $delta = pow($b,2) + (-4* $a * $c);
    if( $delta >= 0) {
        $x1 = (-$b - sqrt($delta)) / 2 * $a;
        $x2 = (-$b + sqrt($delta)) / 2 * $a;
        echo "Δ = $delta <br> x<sub>1</sub> = $x1 <br> x<sub>2</sub> = $x2";
    }else echo "Δ = $delta";
}