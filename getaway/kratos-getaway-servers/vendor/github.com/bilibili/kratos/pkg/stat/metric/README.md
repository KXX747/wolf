统计

  window 
         size
         []bucket
    
    windowRest()->resetBucket()
    Iterator()创建遍历window
    
  bucket
         []points
         count
         next->bucket
     
       resetBucket() 重置bucket
     
   
   

  Iterator  迭代器在窗口内迭代桶。
    	count         int
    	iteratedCount int
    	cur           *Bucket
    	
      Next() Next返回true，所有的桶都已经迭代了。
      Bucket()  桶获得当前下一个桶
      
      
  point_policy 
  
    
    	
    	