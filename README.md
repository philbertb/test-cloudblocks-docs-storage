# Create a new repository on the command line

echo "# cloudblocks-storage-doc" >> README.md
git init
git add README.md
git commit -m "first commit"
git remote add origin git@github.com:thr3z/cloudblocks-storage-doc.git
git push -u origin master

# Push an existing repository from the command line

git remote add origin git@github.com:thr3z/cloudblocks-storage-doc.git
git push -u origin master
