   README GO EBUIILD <!-- /\* Font Definitions \*/ @font-face {font-family:Wingdings; panose-1:5 0 0 0 0 0 0 0 0 0;} @font-face {font-family:"Cambria Math"; panose-1:2 4 5 3 5 4 6 3 2 4;} @font-face {font-family:"Segoe UI Emoji"; panose-1:2 11 5 2 4 2 4 2 2 3;} /\* Style Definitions \*/ p.MsoNormal, li.MsoNormal, div.MsoNormal {margin:0in; font-size:12.0pt; font-family:"Times New Roman",serif;} h1 {mso-style-link:"Heading 1 Char"; margin-right:0in; margin-left:0in; font-size:24.0pt; font-family:"Times New Roman",serif; font-weight:bold;} a:link, span.MsoHyperlink {color:blue; text-decoration:underline;} p {margin-right:0in; margin-left:0in; font-size:12.0pt; font-family:"Times New Roman",serif;} span.Heading1Char {mso-style-name:"Heading 1 Char"; mso-style-link:"Heading 1"; font-family:"Calibri Light",sans-serif; color:#2F5496;} .MsoChpDefault {font-size:10.0pt;} @page WordSection1 {size:8.5in 11.0in; margin:1.0in 1.0in 1.0in 1.0in;} div.WordSection1 {page:WordSection1;} /\* List Definitions \*/ ol {margin-bottom:0in;} ul {margin-bottom:0in;} -->

go-ebuild
=========

* * *

**go-ebuild aided by** "professor ChatGpt PhD.. "

( I dared to asked this repo ULL ...how would I make an ebuild using **_go_**\-**_mod_**.**_elcass_** the Answer was **12 sounds latter...** [vuls.io](vuls.io) metasploit db to vuls.io ... and BAM a proto type ebuild in seconds... how to make a setup.py for this bam... still gobsmacked.. really)

* * *

**Cargo Ebuild** ,,,,,, 💡 **HMMM**

**go-ebuild or similar has been the Tool FOR YEARS People have WANTED, NEED baldly**

_I wont Say it works YET._.  HELP is WANTED/PULLS welcomed..

**' few Dozen more asks for Chatgp**t and have a hopeful prototype.framework . or least getting warmer.. 🔥 **refactor : evaluate :**  <code snip or whole damn thing> and chatgpt helps you even meld functions like a boss..

But so far Chart GPT and the Online Inter have guided me to a possible Prototype heck it even wrote make ebuild function **alone** based on _go-mod.eclass_

**I'm A Beginner to slighly above at best.**  ,  \[Dyslexic Sysadmin/Sec-ops Engineer/ analyst\]   scripts at best programing low hanging of fruit or i dabble  ,  
go-Ebuild  hopefully and or cargo-ebuild are key for any SEC-OPS admin or Pen-testing researcher  , with new tools , and now or having to beg devs bugzilla ROOT 4 ages ,  or fire into repo Yourself and be using em Eh'Voila

... With ChatGpt We punching At Mike Tyson level Fighters today...   **Defiantly Punching WELL above my weight class** but ChatGPT is my M1911 ATM  
 \[first I was like /meh but yeah **GOBSMAKED** at how assistive ChatGPT can be.\]

Colt Guns May have Tamed the west;  kinda moment ChatGPT or alike have that Colt .45 moment in History for sure ...

 but when you need to punch at hypnotic windmills (Don Quixote meme..) (To software devs "the moonbase")  
 and rapidly **Prototype things** **Gregariously** above your skills **chatGPT** can be of great service if you have the patience .

**Stage of Portotype**

*   added a scan main function
*   Chat GPT derived from GO-MOD _writeEbuild_ .eclass  and split off from main that dose final processing of ebuild / egosum$         
*   ADDED get-ego-vendor tool call  >> egosum$  func egosum , sees if installed asks to emerge or build if not on gentoo... Who knows you could be on windows @ work ovr'SSH at lunch
*   main get repo info call         writeEbuild  promt user for git url etc pre-processing
*   CLONE Function git clone FOO /tmp/foo
*   CLEAN UP FOO.git
*   read IN GO go.mod files > func witeEbuild
*   findGoFiles , readin from git clone FOO...  if user says clone repo to /tmp....
*   TO DO

*   findGoFiles likely needs to go-ebuild . {$CWD} also                                                                                                                                                                                               
*   CAN the GOLANG GODS smile upon a helpfull golang PRO to touch this thing down or give me just enough hints to get vesion  00.1  alive ?  
    SPLIT go into FILES so their more bit-sized and easier to manage. and or improve.. do function calls from includes ...

.
