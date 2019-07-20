package main

import "testing"

func TestSum(t *testing.T) {

   	texto := "Code.education Rocks!"
	
   	result := soma()

   if result != texto {
		
	t.Errorf("Texto nao sao semelhantes to be %s but instead got %s!", result, texto)
	
}





}