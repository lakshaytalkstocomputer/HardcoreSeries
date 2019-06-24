import os
import sys



def make_folders(no_of_quest):

	try:
		path = os.getcwd()
	except OSError:
		print("Error in getting path")
	else:
		print("Current working path: "+path)

	name = 'Question'
	folders= ['java','c','cpp','python']
	no_ques = no_of_quest

	for i in range(no_ques):
		level1 = path+"/"+name+str(i+1)
	
		try:
			os.mkdir(level1)
		except OSError:
			print(" Creation of directory failed at level - 1")
		else:
			print(" Succes at level 1")
	
		for j in folders:
			level2= level1+"/"+j
			

			try:
				os.mkdir(level2)
			except OSError:
				print(" Creation of directory failed at level - 2")
			else:
				print(" Succes at level 2")

			f = open(level2+"/STD.in","w")
			f.close()


			g = open(level2+"/STD.out","w")
			g.close()

	
if __name__ == '__main__':
	
	if(len(sys.argv) == 1):
		print("Plese enter no of questions as acommand line argu")
	else:
		q = int(sys.argv[1])
		print(q)
		make_folders(q)