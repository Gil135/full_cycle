packge main

func TestSoma(t * testing.T)  {
	total :=Soma(10 ,15)
	if total != 25{
		t.ErrorF("Resultado")
	}
}