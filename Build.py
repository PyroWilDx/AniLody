import os
import shutil
import subprocess

def buildAniLody():
    targets = [
        ("windows", "amd64", "AniLody-Win64", ".exe"),
        ("darwin", "amd64", "AniLody-Mac64", ""),
        ("linux", "amd64", "AniLody-Linux64", "")
    ]

    distDir = "dist"
    os.makedirs(distDir, exist_ok=True)

    for goos, goarch, outputFolder, ext in targets:
        outputPath = os.path.join(distDir, outputFolder)
        os.makedirs(outputPath, exist_ok=True)

        executableName = f"AniLody{ext}"
        outputExecutable = os.path.join(outputPath, executableName)

        env = os.environ.copy()
        env["GOOS"] = goos
        env["GOARCH"] = goarch

        cmd = ["go", "build", "-o", outputExecutable, "cmd/app/AniLody.go"]
        print(f"Building {goos}/{goarch}...")
        subprocess.run(cmd, env=env, check=True)

        for folder in ["bin", "config"]:
            src = folder
            dest = os.path.join(outputPath, folder)
            if os.path.exists(src):
                if os.path.exists(dest):
                    shutil.rmtree(dest)
                shutil.copytree(src, dest)

    print("Build Complete.")

if __name__ == "__main__":
    buildAniLody()
