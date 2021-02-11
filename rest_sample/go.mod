module rest_sample

go 1.15

replace ponta027.dip.jp/helloworld => ../helloworld

replace ponta027.dip.jp/rest => ./rest

require (
	ponta027.dip.jp/helloworld v0.0.0-00010101000000-000000000000
	ponta027.dip.jp/rest v0.0.0-00010101000000-000000000000 // indirect
)
