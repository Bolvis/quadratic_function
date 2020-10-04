#!/usr/bin/env ruby

$stdout.sync = false
print 'set a ->' 
a = gets.chomp.to_f
print 'set b ->' 
b = gets.chomp.to_f
print 'set c ->' 
c = gets.chomp.to_f

def kwadratowa(a, b, c)
    delta = (b*b) - (4*a*c)
    if delta >= 0 
        x1 = ((-1 * b) - Math.sqrt(delta))/(2*a)
        x2 = ((-1 * b) + Math.sqrt(delta))/(2*a)
        possitiveDelta = true
    end
    return delta, x1, x2, possitiveDelta
end

delta, x1, x2, possitiveDelta = kwadratowa a,b,c

if possitiveDelta
    puts "x1 = " + x1.to_s
    puts "x2 = " + x2.to_s
end
puts "delta = " + delta.to_s