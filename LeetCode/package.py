import os, sys

cur = os.path.abspath('.')
for f in os.listdir(cur):
    print("deal {}".format(f))
    if f.endswith('.py') and not f.startswith('package'):
        dirname = f.split(".")[0].split("_")[0]
        dirname = "{:04}".format(int(dirname))
        dirpath = os.path.join(cur, dirname)
        if not os.path.exists(dirpath):
            os.makedirs(dirpath)
        fpath = os.path.join(dirpath, f)
        
        with open(fpath, 'wb') as fn:
            with open(os.path.join(cur, f), 'rb') as fo:
                fn.write(fo.read())
        os.remove(f)
        # break

