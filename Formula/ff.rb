class Ff < Formula
  desc "A brief description of your application"
  homepage "https://github.com/joshwycuff/homebrew-toolbox"
  url "https://github.com/joshwycuff/homebrew-toolbox/releases/download/ff-f5d40b09850f0cdd99dbc9b8cbcb81d4c2c257a5/ff.tar.gz"
  version "f5d40b09850f0cdd99dbc9b8cbcb81d4c2c257a5"

  depends_on "rg"
  depends_on "bat"

  def install
    bin.install "ff"
    man1.install Dir["*.1"]
    bash_completion.install "ff.bash" => "ff"
    zsh_completion.install "ff.zsh" => "_ff"
    fish_completion.install "ff.fish"
  end
end
