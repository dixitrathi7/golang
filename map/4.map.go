seen := make(map[string]struct{}) 

for _, v := range arr 
{ 
	if _, exists := seen[v]; !exists 
	{ 
		seen[v] = struct{}{} 
		slice = append(slice, v) 
		} 
	}