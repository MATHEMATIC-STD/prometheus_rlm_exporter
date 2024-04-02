package mocks

const STD_OUT_EMPTY_NUKE_R = `
Setting license file path to 4101@RLM-SERVER.domain.domain
rlmutil v15.2
Copyright (C) 2006-2023, Reprise Software, Inc. All rights reserved.


rlm status on RLM-SERVER.domain.domain (port 4101), up 5d 14:26:14
rlm software version v12.2 (build:2)
rlm comm version: v1.2
Startup time: Mon Mar 25 12:29:40 2024
Todays Statistics (22:55:53), init time: Sat Mar 30 04:00:01 2024
Recent Statistics (00:23:38), init time: Sun Mar 31 02:32:16 2024

				Recent Stats         Todays Stats         Total Stats
				00:23:38             22:55:53          5d 14:26:14
Messages:    1462 (1/sec)           64739 (0/sec)          316966 (0/sec)
Connections: 1093 (0/sec)           48291 (0/sec)          243624 (0/sec)

--------- ISV servers ----------
	Name           Port Running Restarts
foundry          49679   Yes      0

------------------------


foundry ISV server status on RLM-SERVER.domain.domain (port 49679), up 5d 14:26:14
foundry software version v12.2 (build: 2)
foundry comm version: v1.2
foundry Debug log filename: C:\ProgramData\The Foundry\RLM\log\\foundry.dlog
foundry Report log filename: <stdout>
Startup time: Mon Mar 25 12:29:41 2024
Todays Statistics (22:55:54), init time: Sat Mar 30 04:00:01 2024
Recent Statistics (00:23:38), init time: Sun Mar 31 02:32:17 2024

				Recent Stats         Todays Stats         Total Stats
				00:23:38             22:55:54          5d 14:26:14
Messages:    1590 (1/sec)           71852 (0/sec)          598788 (1/sec)
Connections: 725 (0/sec)           31843 (0/sec)          170593 (0/sec)
Checkouts:   24 (0/sec)           1333 (0/sec)          61732 (0/sec)
Denials:     0 (0/sec)           0 (0/sec)          1230 (0/sec)
Removals:    0 (0/sec)           0 (0/sec)          358 (0/sec)


------------------------

foundry license pool status on RLM-SERVER.domain.domain (port 49679)

hieroplayer_i v2024.0331, pool: 1
	count: 1, # reservations: 0, inuse: 0, exp: 31-mar-2024
	obsolete: 0, min_remove: 120, total checkouts: 0
nukexassist_i v2024.0331, pool: 2
	count: 2, # reservations: 0, inuse: 0, exp: 31-mar-2024
	obsolete: 0, min_remove: 120, total checkouts: 0
nuke_i v2024.0331, pool: 3
	count: 1, # reservations: 0, inuse: 1, exp: permanent
	obsolete: 0, min_remove: 120, total checkouts: 2321
nukex_i v2024.0331, pool: 4
	count: 1, # reservations: 0, inuse: 1, exp: permanent
	obsolete: 0, min_remove: 120, total checkouts: 346
hieroplayer_i v2024.0701, pool: 5
	count: 22, # reservations: 0, inuse: 0, exp: 1-jul-2024
	obsolete: 0, min_remove: 120, total checkouts: 0
nuke_i v2024.0701, pool: 6
	count: 22, # reservations: 0, inuse: 3, exp: permanent
	obsolete: 0, min_remove: 120, total checkouts: 25854
nukexassist_i v2024.0701, pool: 7
	count: 6, # reservations: 0, inuse: 0, exp: 1-jul-2024
	obsolete: 0, min_remove: 120, total checkouts: 0
nukex_i v2024.0701, pool: 8
	count: 3, # reservations: 0, inuse: 1, exp: permanent
	obsolete: 0, min_remove: 120, total checkouts: 1941
nuke_i v2024.0416, pool: 9
	count: 4, # reservations: 0, inuse: 0, exp: 16-apr-2024
	obsolete: 0, min_remove: 120, total checkouts: 0
nuke_r v2024.0701, pool: 10
	count: 16, # reservations: 0, inuse: 0, exp: permanent
	obsolete: 0, min_remove: 120, total checkouts: 22159
nuke_r v2025.0118, pool: 11
	count: 5, # reservations: 0, inuse: 0, exp: 20-jan-2025
	obsolete: 0, min_remove: 120, total checkouts: 5817
nuke_r v2024.0701, pool: 12
	count: 4, # reservations: 0, inuse: 0, exp: 1-apr-2024
	obsolete: 0, min_remove: 120, total checkouts: 3294


	------------------------

	foundry license usage status on RLM-SERVER.domain.domain (port 49679)

	nuke_i v2024.0331: user1.lastname@mocked-computer-01 1/0 at 03/28 21:44  (handle: 36a5) 
	nuke_i v2024.0331: user1.lastname@mocked-computer-01 1/0 at 03/28 21:44  (handle: 2803) 
	nukex_i v2024.0331: user2.lastname@mocked-computer-02 1/0 at 03/28 20:41  (handle: 374d) 
	nukex_i v2024.0331: user2.lastname@mocked-computer-02 1/0 at 03/28 20:41  (handle: 1ae1) 
	nuke_i v2024.0701: user2.lastname@mocked-computer-02 1/0 at 03/28 20:41  (handle: c0e) 
	nuke_i v2024.0701: user2.lastname@mocked-computer-02 1/0 at 03/28 20:41  (handle: 2ae5) 
	nuke_i v2024.0701: user3.lastname@mocked-computer-03 1/0 at 03/28 18:58  (handle: 3b9d) 
	nuke_i v2024.0701: user4.lastname@mocked-computer-04 1/0 at 03/28 20:39  (handle: 329) 
	nukex_i v2024.0701: user1.lastname@mocked-computer-01 1/0 at 03/28 21:44  (handle: 108f) 
	nukex_i v2024.0701: user1.lastname@mocked-computer-01 1/0 at 03/28 21:44  (handle: e8) 
`

const STD_OUT = `
Setting license file path to 4101@RLM-SERVER.domain.domain
rlmutil v15.2
Copyright (C) 2006-2023, Reprise Software, Inc. All rights reserved.


rlm status on RLM-SERVER.domain.domain (port 4101), up 5d 14:26:14
rlm software version v12.2 (build:2)
rlm comm version: v1.2
Startup time: Mon Mar 25 12:29:40 2024
Todays Statistics (22:55:53), init time: Sat Mar 30 04:00:01 2024
Recent Statistics (00:23:38), init time: Sun Mar 31 02:32:16 2024

				Recent Stats         Todays Stats         Total Stats
				00:23:38             22:55:53          5d 14:26:14
Messages:    1462 (1/sec)           64739 (0/sec)          316966 (0/sec)
Connections: 1093 (0/sec)           48291 (0/sec)          243624 (0/sec)

--------- ISV servers ----------
	Name           Port Running Restarts
foundry          49679   Yes      0

------------------------


foundry ISV server status on RLM-SERVER.domain.domain (port 49679), up 5d 14:26:14
foundry software version v12.2 (build: 2)
foundry comm version: v1.2
foundry Debug log filename: C:\ProgramData\The Foundry\RLM\log\\foundry.dlog
foundry Report log filename: <stdout>
Startup time: Mon Mar 25 12:29:41 2024
Todays Statistics (22:55:54), init time: Sat Mar 30 04:00:01 2024
Recent Statistics (00:23:38), init time: Sun Mar 31 02:32:17 2024

				Recent Stats         Todays Stats         Total Stats
				00:23:38             22:55:54          5d 14:26:14
Messages:    1590 (1/sec)           71852 (0/sec)          598788 (1/sec)
Connections: 725 (0/sec)           31843 (0/sec)          170593 (0/sec)
Checkouts:   24 (0/sec)           1333 (0/sec)          61732 (0/sec)
Denials:     0 (0/sec)           0 (0/sec)          1230 (0/sec)
Removals:    0 (0/sec)           0 (0/sec)          358 (0/sec)


------------------------

foundry license pool status on RLM-SERVER.domain.domain (port 49679)

hieroplayer_i v2024.0331, pool: 1
	count: 1, # reservations: 0, inuse: 0, exp: 31-mar-2024
	obsolete: 0, min_remove: 120, total checkouts: 0
nukexassist_i v2024.0331, pool: 2
	count: 2, # reservations: 0, inuse: 0, exp: 31-mar-2024
	obsolete: 0, min_remove: 120, total checkouts: 0
nuke_i v2024.0331, pool: 3
	count: 1, # reservations: 0, inuse: 1, exp: permanent
	obsolete: 0, min_remove: 120, total checkouts: 2321
nukex_i v2024.0331, pool: 4
	count: 1, # reservations: 0, inuse: 1, exp: permanent
	obsolete: 0, min_remove: 120, total checkouts: 346
hieroplayer_i v2024.0701, pool: 5
	count: 22, # reservations: 0, inuse: 0, exp: 1-jul-2024
	obsolete: 0, min_remove: 120, total checkouts: 0
nuke_i v2024.0701, pool: 6
	count: 22, # reservations: 0, inuse: 3, exp: permanent
	obsolete: 0, min_remove: 120, total checkouts: 25854
nukexassist_i v2024.0701, pool: 7
	count: 6, # reservations: 0, inuse: 0, exp: 1-jul-2024
	obsolete: 0, min_remove: 120, total checkouts: 0
nukex_i v2024.0701, pool: 8
	count: 3, # reservations: 0, inuse: 1, exp: permanent
	obsolete: 0, min_remove: 120, total checkouts: 1941
nuke_i v2024.0416, pool: 9
	count: 4, # reservations: 0, inuse: 0, exp: 16-apr-2024
	obsolete: 0, min_remove: 120, total checkouts: 0
nuke_r v2024.0701, pool: 10
	count: 16, # reservations: 0, inuse: 0, exp: permanent
	obsolete: 0, min_remove: 120, total checkouts: 22159
nuke_r v2025.0118, pool: 11
	count: 5, # reservations: 0, inuse: 0, exp: 20-jan-2025
	obsolete: 0, min_remove: 120, total checkouts: 5817
nuke_r v2024.0701, pool: 12
	count: 4, # reservations: 0, inuse: 0, exp: 1-apr-2024
	obsolete: 0, min_remove: 120, total checkouts: 3294


	------------------------

	foundry license usage status on RLM-SERVER.domain.domain (port 49679)

	nuke_i v2024.0331: user1.lastname@mocked-computer-01 1/0 at 03/28 21:44  (handle: 36a5) 
	nuke_i v2024.0331: user1.lastname@mocked-computer-01 1/0 at 03/28 21:44  (handle: 2803) 
	nukex_i v2024.0331: user2.lastname@mocked-computer-02 1/0 at 03/28 20:41  (handle: 374d) 
	nukex_i v2024.0331: user2.lastname@mocked-computer-02 1/0 at 03/28 20:41  (handle: 1ae1) 
	nuke_i v2024.0701: user2.lastname@mocked-computer-02 1/0 at 03/28 20:41  (handle: c0e) 
	nuke_i v2024.0701: user2.lastname@mocked-computer-02 1/0 at 03/28 20:41  (handle: 2ae5) 
	nuke_i v2024.0701: user3.lastname@mocked-computer-03 1/0 at 03/28 18:58  (handle: 3b9d) 
	nuke_i v2024.0701: user4.lastname@mocked-computer-04 1/0 at 03/28 20:39  (handle: 329) 
	nuke_r v2024.0701: user1.lastname@mocked-computer-05 1/0 at 03/28 20:39  (handle: 329h) 
	nuke_r v2024.0701: user1.lastname@mocked-computer-06 1/0 at 03/28 20:39  (handle: 329s) 
	nukex_i v2024.0701: user1.lastname@mocked-computer-01 1/0 at 03/28 21:44  (handle: 108f) 
	nukex_i v2024.0701: user1.lastname@mocked-computer-01 1/0 at 03/28 21:44  (handle: e8) 
`

const STD_OUT_WRONG = `
Some errors that does not contain any fromatted data.`
