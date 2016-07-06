function result=prime(num)
  count = 0;
  for i=1:num
    if mod(num,i)==0
      count=count+1;
    endif
  endfor
  if count==2
    result=1;
  else
    result=0;
  endif
endfunction
   
function average=percent(limit)
  count=0;
  i=2;
  sum=0;
  while(i>=2)
    if prime(i)==1
      if i<limit
        sum=sum+i;
        count=count+1;
      endif
    endif
    i=i+1;
    if i>limit
      break;
    endif
  endwhile
  average=(count/limit)*100
endfunction
     
function avg=twinpercent(lmt)
  cnt = 0;
  i=2;
  while(i>=2)
    if prime(i)==1
      if prime(i+2)==1
          cnt = cnt+2;
      endif
    endif
    i=i+1;
    if i>=lmt
      break;
    endif
  endwhile
  avg=(cnt/lmt)*100
endfunction
 
 
number = 1:10:1000
      plot(number,arrayfun(@percent, number),'-xr')
      hold on;
      xlabel('Numbers')
      ylabel('Percentage')
      plot(number,arrayfun(@twinpercent, number),'-ob')
      hold on;
      plot(number,log(number),'-og')
      hold off;
      legend('Prime','Twinprimes','logarithm')
