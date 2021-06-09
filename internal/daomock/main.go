package daomock

//go:generate mockgen -destination=mock.go -package=$GOPACKAGE github.com/KennyChenFight/gin-starter/pkg/dao MemberDAO
