package main

import (
	"encoding/base64"
)

func assetMap() map[string][]byte {
	var assets = make(map[string][]byte, 2)

	assets["app.js"], _ = base64.StdEncoding.DecodeString("H4sIAAAJbogA/+xYT3PbNha/61O8KNmQHHMpbbLZ2aGi3ckm29ZtnGQSpxdVB5h8ImGTAAeAXLkaffcOAJKiSND2ZHrpTHmwyIf35/f+AvAtEVDRFJYwX0z0h1REqOMnss5akhOhvqCgKC9IBUtYTQAA9oYeg/fUvPw9qbZeCNLwxbDa03gegoo9IpB4hxD2NP5Hh7AOgZESY5i+/fQVvkqSIfh/C6aH0KldkLKn/WVf+4s+4V+GUFCGXXOf31w05q7uFMpRk5noWfynE/73XPCtogzlqKKkp+iVU9GbouAJUZQz+ExUA29mAa4Xk8kPNMuNRhlJVB8rzSr9fVbwK1LEsN9K/Hr5NoYNKSQeDsFiMnkWZah+/PLxg+/NUqKIF8JmyxIt6evvAPYG8rMISZIbUoeDhpDknCbYsBlW33ta0VR6QUvST0SqClnqP/O919xA+48Xwh5uSbHFuNYDh1MhI6hwp/zaTrAw64f619aoRrWary2p4CT1g8XEuNcgBYkFJurT+Tu/gWpFj2CjDWWp78WWE1MvsIZrSwI3AmVuNHf13rHk/zslsETpt2HQXaFyKt/mtmv0e2Qys5gYBroBHyMlaJahgCfLJXhdTV4Ae5jN4JPAW2QKNojpFUluoOC8agPUSbdJTufb/nQyZQjdLDUozIJB0ALus52wRrs3OypX87UusQZvDfdcAZXAtkUBv+a0QNhWKVGUZQN1+hlX52NUUhYCRiXZhbBlKW4owzS0laurpg5d3ItbWxjd5zBxfzW8h5OMko1C8aULponHadKWOmm/cV4+LlnNAL0gKo8E37LUx2grUVxQ1sFsB+uQiew6TAlnkhcYFTzzp6jxTsMW2LGB2pJ1eNmu1b6dqKwXp0Hd+qcj4r8VTZcenJkOOgPvuXHNUKyTmobM8iBLRyaKfjZcgK+bhZrdBCi87u0mUYEsU/kC6NlZvy61YCo/kBJh2RNb0XWkp+ZiIGD4zt85JeoOdYqYWVELB1HetprfK7jWpWvr0vXAJW3JDvvWueuhc41xy3me7pyI7erqeh3RYd0fxd8RRWDZHRglqUw2Ir0kUUWXtMROplQINy5Itq7UVjBYqRBOVFhkqxbxOvpZz3a5ulmvHV3p6FQ7Eo5OSVRauX/0oifUaebJpOu3QAlLmE4X99aagf9eb8331ZnVpf+edSR0+PULnMH0F9YxdDjZCQueeUGUq7LwBcqg3b26zVihoDylyef7mrLHVDennklPbNs9fw5PkKVd9J1NCwYjwG6ULlN65RHNf09jS1TnTKG4JYXfAx7Ci1fzebCY/MmmgK1Bfbpd/9Xxf1jHHxFf3lXuDB49VkP5OqzVVua+G7u6qzDu2AidXPaI/YBzusjc4low7oTezVUScaPPLG6gYLZ/clVgWh/RnWyHB+PaGYVgppBr2xpiqC8lbnAlERll73GjYvj3PNTnnZ8QKyBFYQUlFLhRQAqaMUydOmRFEsqyS17F8GLuDlHN8z+uFC/H2fS569Ik1tt5bhbCaGkuS3Uwh1yHIUlRVeBYDPR1IK5nzYhN7X0Mng7FCCwbyBhGHNvF8HL+GKSJwJQqOYb1gUJyKCwwQ5Z+u74BzZzsR2Npk5cShYqWOBKsRHApc0JFDEpsR4JuTt2jhsDuRs1pPj65sbkV6qd/C4gHlMc259DG3b2Bua8Coa1Cfcn6dgSK80LR6nHJdgdpNqs4Zeo7LkqiYvD25jO6O4ykcjbLkaQoWgEHnwNpVfDmHxljaJt/nIyH7OFRAI+NXGPM/g7XcVdxoa+839pIp5M86B8r7dHxdwAAAP//AQAA//8M58QpnxMAAA==")
	assets["index.html"], _ = base64.StdEncoding.DecodeString("H4sIAAAJbogA/7xUP2/bPhTc9Sn44/TrILF20AyupKXukKkeunSkyWeRKUWyfFRqoeh3L0iptaEISICm8WK+P7o7nqCr/9t/+vD5y+EjUbE3bVGnP2K47RoKlrZFUSvgsi0IIaQ22n4lAUxDMY4GUAFESlSAU0NVjB53jPX8LKStjs5FjIH7VAjXsz8NdlPdVLdMIF56Va9tJRApEcEhuqA7bRvKrbNj7waks4BMS+LooaERzjGhzLNKKB4i+ZGL9FOgOxV3ZPvurT+/z+2fEwrLMG1Rs+luRX10cpwppH4gwnDEhgpnI9cWQnkyg5Yz0XIruO9Xk8cYpuxlebtcObnQ/95J58V8spsfwZCTCw31WiJtD3f7Xc1ye2UdwYCIRMt5nTgrFLcdNHQaHe72/79ZI2LTfCGRJV3LntQPVzYsynTvRJ9fBb6kX4/WEgO9kJXCD7Rd6FmR+FJ0gffPoVsz6J/a0IXXdKETr2HCZrumywfISozrkggf4Lki5uOcKCJoHwkGcYkw4SRU998GCGOOrulYbqrNttrmqLrHxDk92z4BpHSnpg8ig12Vf4HSOzkYQAZn70LUtnsKjHu/WKnZlHopBnP2/wIAAP//AQAA//9TiF2CDAYAAA==")
	return assets
}
